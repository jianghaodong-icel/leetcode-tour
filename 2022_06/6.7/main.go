package main

import "github.com/emirpasic/gods/trees/redblacktree"

type pair struct {
	score int
	name  string
}

func compare(x, y interface{}) int {
	a, b := x.(pair), y.(pair)
	if a.score > b.score || (a.score == b.score && a.name < b.name) {
		return -1
	}
	return 1
}

type SORTracker struct {
	*redblacktree.Tree
	cur redblacktree.Iterator
}

func Constructor() SORTracker {
	root := redblacktree.NewWith(compare)
	root.Put(pair{}, nil)
	return SORTracker{root, root.IteratorAt(root.Left())}
}

func (this *SORTracker) Add(name string, score int) {
	p := pair{score: score, name: name}
	this.Put(p, nil)
	if compare(p, this.cur.Key()) < 0 {
		this.cur.Prev()
	}
}

func (this *SORTracker) Get() string {
	name := this.cur.Key().(pair).name
	this.cur.Next()
	return name
}

/**
 * Your SORTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(name,score);
 * param_2 := obj.Get();
 */

func main() {

}
