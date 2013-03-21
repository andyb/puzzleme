//app code
$(function() {
    $(".drag").draggable({revert: "valid", snap:true, stack: ".drag"});
    $(".drag").droppable({
      drop: function( event, ui ) {
        $( this )            
            var dropImg = ui.draggable.context.innerHTML;               
            ui.draggable.html(this.innerHTML);
            this.innerHTML = dropImg;
                        
      }
    });
  
});