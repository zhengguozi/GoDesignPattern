package Memento

// 备份类
type Memento struct {
	State string
}

func (m *Memento) GetSavedState() string {
	return m.State
}

// -------------------------------------------------------
// 需要备份的原始类
type Originator struct {
	State string
}

func (e *Originator) SetState(state string) {
	e.State = state
}

func (e *Originator) GetState() string {
	return e.State
}

// 怎么备份，怎么恢复（从备份类）
func (e *Originator) CreateMemento() *Memento {
	return &Memento{e.State}
}
func (e *Originator) RestoreMemento(m *Memento) {
	e.SetState(m.GetSavedState())
}

// -------------------------------------------------------
// 负责人类 维护所有备份（栈）
type Caretaker struct {
	MementoArray []*Memento
}

func (c *Caretaker) AddMemento(m *Memento) {
	c.MementoArray = append(c.MementoArray, m)
}

func (c *Caretaker) GetMemento(index int) *Memento {
	return c.MementoArray[index]
}
