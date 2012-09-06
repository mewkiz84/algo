package uf

type Uf []int

func New(n int) Uf {
   u := make(Uf, n)
   for key := range u {
      u[key] = key
   }
   return u
}

func (set Uf) Union(p, q int) {
   i := set.Root(p)
   j := set.Root(q)
   set[i] = j
}

func (set Uf) Connected(p, q int) bool {
   return set.Root(p) == set.Root(q)
}

func (set Uf) Root(p int) int {
   for p != set[p] {
      p = set[p]
   }
   return p
}
