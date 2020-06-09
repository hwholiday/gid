package service

import (
	"context"
	"gid/entity"
	"sync"
	"time"
)

type Alloc struct {
	Mu        sync.RWMutex
	BizTagMap map[string]*BizAlloc
}

type BizAlloc struct {
	Mu      sync.Mutex
	Wg      sync.WaitGroup
	BazTag  string
	IdArray []*IdArray
	GetDb   bool //当前正在查询DB
}

type IdArray struct {
	Cur   int64 //当前发到哪个位置
	Start int64 //最小值
	End   int64 //最大值
}

func (s *Service) NewAllocId() (a *Alloc, err error) {
	var res []entity.Segments
	if res, err = s.r.SegmentsGetAll(); err != nil {
		return
	}
	a = &Alloc{
		BizTagMap: make(map[string]*BizAlloc),
	}
	for _, v := range res {
		a.BizTagMap[v.BizTag] = &BizAlloc{
			BazTag:  v.BizTag,
			GetDb:   false,
			IdArray: make([]*IdArray, 0),
		}
	}
	return
}

func (b *BizAlloc) GetId() (id int64, err error) {
	var (
		canGetId    bool
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	)
	defer cancel()
	b.Mu.Lock()
	if b.LeftIdCount() > 0 {
		id = b.PopId()
		canGetId = true
	}
	// 2, 段<=1个, 启动补偿线程
	if len(b.IdArray) <= 1 && !b.GetDb {
		b.GetDb = true
		b.Wg.Add(1)
		//go bizAlloc.fillSegments()
	}
	b.Mu.Unlock()
	if canGetId {
		return
	}
	select {
	case <-ctx.Done(): //执行超时
		b.Wg.Done()
	}
	b.Wg.Wait()

	//检查是否需要补充数据

}

func (b *BizAlloc) LeftIdCount() (count int64) {
	for _, v := range b.IdArray {
		arr := v
		//结束位置-开始位置-已经分配的位置
		count += arr.End - arr.Start - arr.Cur
	}
	return count
}

func (b *BizAlloc) PopId() (id int64) {
	id = b.IdArray[0].Start + b.IdArray[0].Cur //开始位置加上分配次数
	b.IdArray[0].Cur++                         //分配次数 +1
	if id+1 >= b.IdArray[0].End {              //该数组里面没有ID了
		b.IdArray = append(b.IdArray[:0], b.IdArray[1:]...) //把分配完的数组移除
	}
	return
}
