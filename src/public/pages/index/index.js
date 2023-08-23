function onRemoveTask(id) {
  $.ajax({
    url: "/remove-task",
    method: "GET",
    data: {id},
    success: function(response) {
        $(".response").html(response); 
    },
    error: function(error) {
        alert("Ошибка при отправке запроса:", error);
    }
  });
}

// Получаем все кнопки для разворачивания/сворачивания подзадач
const toggleSubtaskButtons = document.querySelectorAll('.toggle-subtasks');

// Обработчик события для клика на кнопку
toggleSubtaskButtons.forEach(button => {
    button.addEventListener('click', function(event) {
        const taskId = this.getAttribute('data-task-id');
        const subtasks = document.getElementById(`subtasks-${taskId}`);
        
        if (subtasks.style.display === 'none') {
            subtasks.style.display = 'block';
        } else {
            subtasks.style.display = 'none';
        }
    });
});
