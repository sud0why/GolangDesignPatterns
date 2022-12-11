package main

import "fmt"

//访问者模式
//操作一组（继承相同父类或实现相同接口的）对象
//难以理解，尽量不要使用

//公司里很多员工，领导根据员工KPI决定薪酬发放，这个功能不应该实现在员工类里

// Employee 员工接口
type Employee interface {
	Score() int
	Accept(visitor EmployeeVisitor) // 接受访问者对象
}

type PM struct {
	NumOfProducts int
}

func (p *PM) Score() int {
	return p.NumOfProducts
}

func (p *PM) Accept(visitor EmployeeVisitor) {
	visitor.VisitPM(p)
}

type QA struct {
	NumOfBugs int
}

func (q *QA) Score() int {
	return q.NumOfBugs
}

func (q *QA) Accept(visitor EmployeeVisitor) {
	visitor.VisitQA(q)
}

type RD struct {
	NumOfNeeds int
}

func (r *RD) Score() int {
	return r.NumOfNeeds
}

func (r *RD) Accept(visitor EmployeeVisitor) {
	visitor.VisitRD(r)
}

// EmployeeVisitor 员工访问者接口
type EmployeeVisitor interface {
	VisitPM(*PM)
	VisitQA(*QA)
	VisitRD(*RD)
}

type Leader struct {
}

func (l Leader) VisitPM(pm *PM) {
	fmt.Printf("Products: %d\n", pm.Score())
}

func (l Leader) VisitQA(qa *QA) {
	fmt.Printf("Bugs: %d\n", qa.Score())
}

func (l Leader) VisitRD(rd *RD) {
	fmt.Printf("Needs: %d\n", rd.Score())
}

func main() {
	rd := RD{NumOfNeeds: 100}
	qa := QA{NumOfBugs: 90}
	pm := PM{NumOfProducts: 80}
	ld := Leader{}
	ld.VisitRD(&rd)
	ld.VisitQA(&qa)
	ld.VisitPM(&pm)
}
