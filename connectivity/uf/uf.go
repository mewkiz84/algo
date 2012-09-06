package uf

type Component struct {
   Number int
   Group int
   Next int
   Size int
}

type Uf []Component

func New(n int) Uf {
   u := make([]Component, n)
   for key := range u {
      u[key].Number = key
      u[key].Group = key
      u[key].Next = key
      u[key].Size = 1
   }
   return u
}

// add connection between p and q
func (union Uf) Union(p, q int) {
   qg := union[q].Group
   pg := union[p].Group
   if pg == qg {
      return
   } else if union[pg].Size < union[qg].Size {
      // add pg to qg and update all number groups in pg to be the group of qg
      union[qg].Size += union[pg].Size
      union[pg].Size = 0
      union.setGroup(pg, qg)
   } else {
      // add qg to pg and update all number groups in qg to be the group of pg
      union[pg].Size += union[qg].Size
      union[qg].Size = 0
      union.setGroup(qg, pg)
   }
   // connect the linked list of p and q
   union[p].Next, union[q].Next = union[q].Next, union[p].Next
}

// return the group root
func (union Uf) Root(p int) int {
   return union[p].Group
}

// set p group to q group
func (union Uf) setGroup(p, q int) {
   for tmp := union[p].Next; union[p].Group != p; tmp = union[tmp].Next {
      union[tmp].Group = q
   }
}

// are p and q connected?
func (union Uf) Connected(p, q int) bool {
   return union[p].Group == union[q].Group
}

// find what component p is in
func (union Uf) Find(p int) int {
   return union[p].Group
}

// return the number of components
func (union Uf) Count() (c int) {
   for _, obj := range union {
      if obj.Size > 0 {
         c++
      }
   }
   return c
}






//type Component struct {
//   N []int
//   G int
//}
//
//type Uf []Component
//
//func New(n int) Uf {
//   return make([]Component, n)
//}
//
//// add connection between p and q
//func (union Uf) Union(p, q int) {
//   qg := union[q].G
//   pg := union[p].G
//   if pg == qg {
//      return
//   } else if len(union[pg].N) >= len(union[qg].N) {
//      union[pg].N = append(union[pg].N, union[qg].N...)
//      for _, obj := range union[qg].N {
//         union[obj].G = union[pg].G
//      }
//      union[qg].N = nil
//   } else {
//      union[qg].N = append(union[qg].N, union[pg].N...)
//      for _, obj := range union[pg].N {
//         union[obj].G = union[qg].G
//      }
//      union[pg].N = nil
//   }
//}
//
//// are p and q connected?
//func (union Uf) Connected(p, q int) bool {
//   return union[p].G == union[q].G
//}
//
//// find what component p is in
//func (union Uf) Find(p int) int {
//   return union[p].G
//}
//
//// return the number of components
//func (union Uf) Count() (c int) {
//   for _, obj := range union {
//      if obj.N != nil {
//         c++
//      }
//   }
//   return c
//}
