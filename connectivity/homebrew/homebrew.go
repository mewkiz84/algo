package homebrew

type Component struct {
   Number int
   Group int
   Next int
   Size int
}

type Homebrew []Component

func New(n int) Homebrew {
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
func (union Homebrew) Union(p, q int) {
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
func (union Homebrew) Root(p int) int {
   return union[p].Group
}

// set p group to q group
func (union Homebrew) setGroup(p, q int) {
   for tmp := union[p].Next; union[p].Group != p; tmp = union[tmp].Next {
      union[tmp].Group = q
   }
}

// are p and q connected?
func (union Homebrew) Connected(p, q int) bool {
   return union[p].Group == union[q].Group
}

// find what component p is in
func (union Homebrew) Find(p int) int {
   return union[p].Group
}

// return the number of components
func (union Homebrew) Count() (c int) {
   for _, obj := range union {
      if obj.Size > 0 {
         c++
      }
   }
   return c
}
