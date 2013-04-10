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


function main($scope) {    
     $scope.puzzleSlices = [
        {img:'temp7'},
        {img:'temp1'},
        {img:'temp8'},
        {img:'temp3'},
        {img:'temp9'},
        {img:'temp5'},
        {img:'temp2'},
        {img:'temp4'},
        {img:'temp6'}
     ]

     $scope.peek = function() {
         $("#puzzle").hide()
         $("#peek").show()
     }

     $scope.puzzle = function() {
         $("#puzzle").show()
         $("#peek").hide()
     }
}