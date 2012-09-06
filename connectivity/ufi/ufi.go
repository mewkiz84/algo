package ufi

type Ufi []ufi

type ufi struct {
   Number int
   Size int
}

func New(n int) Ufi {
   set := make(Ufi, n)
   for key := range set {
      set[key].Number = key
   }
   return set
}

func (set Ufi) Connected(p, q int) bool {
   return set[p] == set[q]
}

func (set Ufi) Union(p, q int) {
   pg := set.Root(p)
   qg := set.Root(q)
   if pg == qg {
      return
   } else if set[pg].Size < set[qg].Size {
      set[pg].Number = qg
      set[qg].Size += set[pg].Size + 1
   } else {
      set[qg].Number = pg
      set[pg].Size += set[qg].Size + 1
   }
}

func (set Ufi) Root(p int) int {
   for p != set[p].Number {
      set[p].Number = set[set[p].Number].Number
      p = set[p].Number
   }
   return p
}
