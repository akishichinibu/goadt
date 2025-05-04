package goadt

type Tuple2[T1 any, T2 any] struct {
	t1 T1
	t2 T2
}

func (t *Tuple2[T1, T2]) Get1() T1 {
	return t.t1
}
func (t *Tuple2[T1, T2]) Get2() T2 {
	return t.t2
}
func (t *Tuple2[T1, T2]) Unwrap() (T1, T2) {
	return t.t1, t.t2
}
func NewTuple2[T1 any, T2 any](t1 T1, t2 T2) Tuple2[T1, T2] {
	return Tuple2[T1, T2]{t1: t1, t2: t2}
}

type Tuple3[T1 any, T2 any, T3 any] struct {
	t1 T1
	t2 T2
	t3 T3
}

func (t *Tuple3[T1, T2, T3]) Get1() T1 {
	return t.t1
}
func (t *Tuple3[T1, T2, T3]) Get2() T2 {
	return t.t2
}
func (t *Tuple3[T1, T2, T3]) Get3() T3 {
	return t.t3
}
func (t *Tuple3[T1, T2, T3]) Unwrap() (T1, T2, T3) {
	return t.t1, t.t2, t.t3
}
func NewTuple3[T1 any, T2 any, T3 any](t1 T1, t2 T2, t3 T3) Tuple3[T1, T2, T3] {
	return Tuple3[T1, T2, T3]{t1: t1, t2: t2, t3: t3}
}

type Tuple4[T1 any, T2 any, T3 any, T4 any] struct {
	t1 T1
	t2 T2
	t3 T3
	t4 T4
}

func (t *Tuple4[T1, T2, T3, T4]) Get1() T1 {
	return t.t1
}
func (t *Tuple4[T1, T2, T3, T4]) Get2() T2 {
	return t.t2
}
func (t *Tuple4[T1, T2, T3, T4]) Get3() T3 {
	return t.t3
}
func (t *Tuple4[T1, T2, T3, T4]) Get4() T4 {
	return t.t4
}
func (t *Tuple4[T1, T2, T3, T4]) Unwrap() (T1, T2, T3, T4) {
	return t.t1, t.t2, t.t3, t.t4
}
func NewTuple4[T1 any, T2 any, T3 any, T4 any](t1 T1, t2 T2, t3 T3, t4 T4) Tuple4[T1, T2, T3, T4] {
	return Tuple4[T1, T2, T3, T4]{t1: t1, t2: t2, t3: t3, t4: t4}
}

type Tuple5[T1 any, T2 any, T3 any, T4 any, T5 any] struct {
	t1 T1
	t2 T2
	t3 T3
	t4 T4
	t5 T5
}

func (t *Tuple5[T1, T2, T3, T4, T5]) Get1() T1 {
	return t.t1
}
func (t *Tuple5[T1, T2, T3, T4, T5]) Get2() T2 {
	return t.t2
}
func (t *Tuple5[T1, T2, T3, T4, T5]) Get3() T3 {
	return t.t3
}
func (t *Tuple5[T1, T2, T3, T4, T5]) Get4() T4 {
	return t.t4
}
func (t *Tuple5[T1, T2, T3, T4, T5]) Get5() T5 {
	return t.t5
}
func (t *Tuple5[T1, T2, T3, T4, T5]) Unwrap() (T1, T2, T3, T4, T5) {
	return t.t1, t.t2, t.t3, t.t4, t.t5
}
func NewTuple5[T1 any, T2 any, T3 any, T4 any, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5) Tuple5[T1, T2, T3, T4, T5] {
	return Tuple5[T1, T2, T3, T4, T5]{t1: t1, t2: t2, t3: t3, t4: t4, t5: t5}
}

type Tuple6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any] struct {
	t1 T1
	t2 T2
	t3 T3
	t4 T4
	t5 T5
	t6 T6
}

func (t *Tuple6[T1, T2, T3, T4, T5, T6]) Get1() T1 {
	return t.t1
}
func (t *Tuple6[T1, T2, T3, T4, T5, T6]) Get2() T2 {
	return t.t2
}
func (t *Tuple6[T1, T2, T3, T4, T5, T6]) Get3() T3 {
	return t.t3
}
func (t *Tuple6[T1, T2, T3, T4, T5, T6]) Get4() T4 {
	return t.t4
}
func (t *Tuple6[T1, T2, T3, T4, T5, T6]) Get5() T5 {
	return t.t5
}
func (t *Tuple6[T1, T2, T3, T4, T5, T6]) Get6() T6 {
	return t.t6
}
func (t *Tuple6[T1, T2, T3, T4, T5, T6]) Unwrap() (T1, T2, T3, T4, T5, T6) {
	return t.t1, t.t2, t.t3, t.t4, t.t5, t.t6
}
func NewTuple6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6) Tuple6[T1, T2, T3, T4, T5, T6] {
	return Tuple6[T1, T2, T3, T4, T5, T6]{t1: t1, t2: t2, t3: t3, t4: t4, t5: t5, t6: t6}
}

type Tuple7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any] struct {
	t1 T1
	t2 T2
	t3 T3
	t4 T4
	t5 T5
	t6 T6
	t7 T7
}

func (t *Tuple7[T1, T2, T3, T4, T5, T6, T7]) Get1() T1 {
	return t.t1
}
func (t *Tuple7[T1, T2, T3, T4, T5, T6, T7]) Get2() T2 {
	return t.t2
}
func (t *Tuple7[T1, T2, T3, T4, T5, T6, T7]) Get3() T3 {
	return t.t3
}
func (t *Tuple7[T1, T2, T3, T4, T5, T6, T7]) Get4() T4 {
	return t.t4
}
func (t *Tuple7[T1, T2, T3, T4, T5, T6, T7]) Get5() T5 {
	return t.t5
}
func (t *Tuple7[T1, T2, T3, T4, T5, T6, T7]) Get6() T6 {
	return t.t6
}
func (t *Tuple7[T1, T2, T3, T4, T5, T6, T7]) Get7() T7 {
	return t.t7
}
func (t *Tuple7[T1, T2, T3, T4, T5, T6, T7]) Unwrap() (T1, T2, T3, T4, T5, T6, T7) {
	return t.t1, t.t2, t.t3, t.t4, t.t5, t.t6, t.t7
}
func NewTuple7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7) Tuple7[T1, T2, T3, T4, T5, T6, T7] {
	return Tuple7[T1, T2, T3, T4, T5, T6, T7]{t1: t1, t2: t2, t3: t3, t4: t4, t5: t5, t6: t6, t7: t7}
}

type Tuple8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any] struct {
	t1 T1
	t2 T2
	t3 T3
	t4 T4
	t5 T5
	t6 T6
	t7 T7
	t8 T8
}

func (t *Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Get1() T1 {
	return t.t1
}
func (t *Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Get2() T2 {
	return t.t2
}
func (t *Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Get3() T3 {
	return t.t3
}
func (t *Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Get4() T4 {
	return t.t4
}
func (t *Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Get5() T5 {
	return t.t5
}
func (t *Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Get6() T6 {
	return t.t6
}
func (t *Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Get7() T7 {
	return t.t7
}
func (t *Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Get8() T8 {
	return t.t8
}
func (t *Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Unwrap() (T1, T2, T3, T4, T5, T6, T7, T8) {
	return t.t1, t.t2, t.t3, t.t4, t.t5, t.t6, t.t7, t.t8
}
func NewTuple8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8) Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{t1: t1, t2: t2, t3: t3, t4: t4, t5: t5, t6: t6, t7: t7, t8: t8}
}

type Tuple9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any] struct {
	t1 T1
	t2 T2
	t3 T3
	t4 T4
	t5 T5
	t6 T6
	t7 T7
	t8 T8
	t9 T9
}

func (t *Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Get1() T1 {
	return t.t1
}
func (t *Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Get2() T2 {
	return t.t2
}
func (t *Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Get3() T3 {
	return t.t3
}
func (t *Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Get4() T4 {
	return t.t4
}
func (t *Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Get5() T5 {
	return t.t5
}
func (t *Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Get6() T6 {
	return t.t6
}
func (t *Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Get7() T7 {
	return t.t7
}
func (t *Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Get8() T8 {
	return t.t8
}
func (t *Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Get9() T9 {
	return t.t9
}
func (t *Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Unwrap() (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	return t.t1, t.t2, t.t3, t.t4, t.t5, t.t6, t.t7, t.t8, t.t9
}
func NewTuple9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9) Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{t1: t1, t2: t2, t3: t3, t4: t4, t5: t5, t6: t6, t7: t7, t8: t8, t9: t9}
}

type Tuple10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any] struct {
	t1  T1
	t2  T2
	t3  T3
	t4  T4
	t5  T5
	t6  T6
	t7  T7
	t8  T8
	t9  T9
	t10 T10
}

func (t *Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get1() T1 {
	return t.t1
}
func (t *Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get2() T2 {
	return t.t2
}
func (t *Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get3() T3 {
	return t.t3
}
func (t *Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get4() T4 {
	return t.t4
}
func (t *Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get5() T5 {
	return t.t5
}
func (t *Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get6() T6 {
	return t.t6
}
func (t *Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get7() T7 {
	return t.t7
}
func (t *Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get8() T8 {
	return t.t8
}
func (t *Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get9() T9 {
	return t.t9
}
func (t *Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Get10() T10 {
	return t.t10
}
func (t *Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Unwrap() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) {
	return t.t1, t.t2, t.t3, t.t4, t.t5, t.t6, t.t7, t.t8, t.t9, t.t10
}
func NewTuple10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10) Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{t1: t1, t2: t2, t3: t3, t4: t4, t5: t5, t6: t6, t7: t7, t8: t8, t9: t9, t10: t10}
}
