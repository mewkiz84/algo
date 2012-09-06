package qf

type Qf []int

func New(n int) Qf {
   obj := make(Qf, n)
   for key := range obj {
      obj[key] = key
   }
   return obj
}

func (set Qf) Connected(p, q int) bool {
   return set[p] == set[q]
}

func (set Qf) Union(p, q int) {
   pg := set[p]
   qg := set[q]
   if pg == qg {
      return
   }
   for key := range set {
      if set[key] == pg {
         set[key] = qg
      }
   }
}
