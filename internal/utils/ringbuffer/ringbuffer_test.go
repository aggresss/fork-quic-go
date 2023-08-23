package ringbuffer

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RingBuffer", func() {
	It("push and pop", func() {
		r := RingBuffer[int]{}
		Expect(len(r.ring)).To(Equal(0))
		Expect(func() { r.PopFront() }).To(Panic())
		r.PushBack(1)
		r.PushBack(2)
		r.PushBack(3)
		Expect(r.PopFront()).To(Equal(1))
		Expect(r.PopFront()).To(Equal(2))
		r.PushBack(4)
		r.PushBack(5)
		Expect(r.Len()).To(Equal(3))
		r.PushBack(6)
		Expect(r.Len()).To(Equal(4))
		Expect(r.PopFront()).To(Equal(3))
		Expect(r.PopFront()).To(Equal(4))
		Expect(r.PopFront()).To(Equal(5))
		Expect(r.PopFront()).To(Equal(6))
	})
	It("clear", func() {
		r := RingBuffer[int]{}
		r.Init(2)
		r.PushBack(1)
		r.PushBack(2)
		Expect(r.full).To(BeTrue())
		r.Clear()
		Expect(r.full).To(BeFalse())
		Expect(r.Len()).To(Equal(0))
	})
	It("front, back, offset", func() {
		r := RingBuffer[int]{}
		r.Init(3)
		r.PushBack(1)
		r.PushBack(2)
		r.PushBack(3)
		r.PopFront()
		r.PushBack(4)
		Expect(r.full).To(BeTrue())
		Expect(*r.Front()).To(Equal(2))
		Expect(*r.Back()).To(Equal(4))
		Expect(*r.Offset(0)).To(Equal(2))
		Expect(*r.Offset(1)).To(Equal(3))
		Expect(*r.Offset(2)).To(Equal(4))
		*r.Front() = 4
		*r.Back() = 2
		Expect(*r.Offset(0)).To(Equal(4))
		Expect(*r.Offset(1)).To(Equal(3))
		Expect(*r.Offset(2)).To(Equal(2))
	})
})
