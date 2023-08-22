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