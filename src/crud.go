package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//АХТУНГ: ФУНКЦИЯ УДАЛЕНИЯ ЗАДАЧ ТОЛЬКО ДЛЯ АДМИНА
/*
TODO:
Добавить проверку, за админа ли удаляем
*/
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	argStr := r.URL.Path[len("/remove-task/"):]
	index, err := strconv.Atoi(argStr)
	if err != nil || index < 0 || index >= len(Tasks) {
		http.Error(w, "Такого индекса нет", http.StatusBadRequest)
		return
	}
	parent := Tasks[index].Parent
	if len(Tasks[index].Children) > 0 {
		http.Error(w, "Задачу с имеющимися подзадачами удалить нельзя", http.StatusBadRequest)
		return
	}
	if parent != nil {
		// Найден ребенок у родителя, удаляем его из списка детей
		for _, child := range parent.Children {
			if child.ID == index {
				// Найден ребенок у родителя, удаляем его из списка детей
				parent.Children, err = removeElement(parent.Children, child)
				if err != nil {
					http.Error(w, "Ошибка при удалении ребенка у родителя", http.StatusInternalServerError)
					return
				}
				break
			}
		}
	}
	// У детей перемещаем к родителю повыше
	/*for _, child := range Tasks[index].Children {
		child.Parent = parent
	}
	*/

	Tasks = append(Tasks[:index], Tasks[index+1:]...)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func removeElement(slice []*Task, element *Task) ([]*Task, error) {
	indexToRemove := -1
	for i, obj := range slice {
		if obj.ID == element.ID {
			indexToRemove = i
			break
		}
	}

	if indexToRemove == -1 {
		return nil, fmt.Errorf("Объект не найден")
	}

	newSlice := append(slice[:indexToRemove], slice[indexToRemove+1:]...)
	return newSlice, nil
}

func CompleteTask(w http.ResponseWriter, r *http.Request) {
	argStr := r.URL.Path[len("/remove-task/"):]
	index, err := strconv.Atoi(argStr)
	if err != nil || index < 0 || index >= len(Tasks) {
		http.Error(w, "Такого индекса нет", http.StatusBadRequest)
		return
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	indexStr := r.URL.Path[len("/update-task/"):]
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(Tasks) {
		http.Error(w, "Такого индекса нет", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	description := r.FormValue("description")
	priorityStr := r.FormValue("priority")
	priority, err := strconv.Atoi(priorityStr)
	if err != nil {
		http.Error(w, "Некорректный приоритет", http.StatusBadRequest)
		return
	}

	backgroundStr := r.FormValue("background")
	background := backgroundStr == "on"

	notificationTime := r.FormValue("notification_time")

	parent := Tasks[index].Parent
	children := Tasks[index].Children

	fmt.Println("Фишки Работают!")
	Tasks[index] = Task{
		ID:               index,
		Title:            title,
		Description:      description,
		Priority:         priority,
		Background:       background,
		NotificationTime: notificationTime,
		Parent:           parent,
		Children:         children,
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
