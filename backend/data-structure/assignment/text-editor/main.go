package main

import (
	"fmt"

	"github.com/ruang-guru/playground/backend/data-structure/assignment/text-editor/stack"
)

// Implementasikan teks editor sederhana
// Teks editor ini memiliki 4 method
// Write: digunakan untuk menulis per karakter
// Read: digunakan untuk mencetak karakter yang telah ditulis
// Undo: digunakan untuk melakukan operasi undo
// Redo: digunakan untuk melakukan operasi redo

type TextEditor struct {
	UndoStack *stack.Stack // membatalkan perintah terakhir
	RedoStack *stack.Stack // memanggil perintah terakhir
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		UndoStack: stack.NewStack(),
		RedoStack: stack.NewStack(),
	}
}

func (te *TextEditor) Write(ch rune) {
	// TODO: answer here
	// call func setToEmpty :: mengosongkan stack (RedoStack)
	te.RedoStack.SetToEmpty()
	// call func Push (UndoStack)
	te.UndoStack.Push(ch)
}

func (te *TextEditor) Read() []rune {
	// TODO: answer here
	// initiate a stack (data temporary) -> call stack.NewStack()
	tmp := stack.NewStack()
	// sediakan variabel buat return []rune
	var r []rune
	// looping (selama te.UndoStack.Top > -1)
	for te.UndoStack.Top > -1 {
		// -> call func UndoStack.Pop
		ch, _ := te.UndoStack.Pop()
			// -> push ke temp stack (temporary.Push(..))
		tmp.Push(ch)
	}
	// looping temporary stack (selama Top > -1)
	for tmp.Top > -1 {
		// -> call func Peek -> kembalian rune
		ch,_ := tmp.Peek()
		// -> munculin output string kembalian value peek (fmt, log)
		fmt.Println(string(ch))
		// -> append rune dari Peek ke return variable
		r = append(r,ch)
		// -> call func te Push(rune dari Peek)
		te.UndoStack.Push(ch)
		// -> call func Pop temporary stack
		tmp.Pop()
	}
	fmt.Println(string(r))
	
	return r

	// return result
}

func (te *TextEditor) Undo() {
	// TODO: answer here
	// call func UndoStack Peek :: ambil element terakhir dri UndoStack
	ch, err := te.UndoStack.Peek()
	// -> jgn lupa check error -> jika ada error return
	if err != nil {
		return
	}
	// call func UndoStack Pop :: hapus element terakhir dri UndoStack
	te.UndoStack.Pop()
	// call func RedoStack Push(hasil Peek UndoStack) :: tambah element terakhir yg diundo ke RedoStack
	te.RedoStack.Push(ch)
}

func (te *TextEditor) Redo() {
	// TODO: answer here
	ch, err := te.RedoStack.Peek()
	// -> jgn lupa check error -> jika ada error return
	if err != nil {
		return
	}
	// call func UndoStack Pop :: hapus element terakhir dri UndoStack
	te.RedoStack.Pop()
	// call func RedoStack Push(hasil Peek UndoStack) :: tambah element terakhir yg diundo ke RedoStack
	te.UndoStack.Push(ch)
	// call func RedoStack Peek :: ambil element terakhir dri RedoStack
	// -> check error
	// call func RedoStack Pop :: hapus element terakhir dri RedoStack
	// call func UndoStack Push(hasil Peek RedoStack) :: tambah element terakhir yg diredo ke UndoStack
}

func main() {
	te := NewTextEditor()
	te.Write('h')
	te.Write('e')
	te.Write('l')
	te.Write('l')
	te.Write('o')
	te.Undo()
	te.Read()
	te.Redo()
	te.Read()

}
