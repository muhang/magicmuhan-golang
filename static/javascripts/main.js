$(document).ready(function () {

  var workSelectedClass = 'work-selected';
  var workInactiveClass = 'work-inactive';
  var workDefaultClass = 'work-default';

  function animateSkills() {
    if ($('#about-wrapper').length) {
      $('#about-wrapper').addClass('show-section');    
    }
  }

  function clearWorkSelection() {
    $('#work-section-jobs').removeClass(workSelectedClass);
    $('#work-section-projects').removeClass(workSelectedClass);
    $('#work-section-jobs').removeClass(workInactiveClass);
    $('#work-section-projects').removeClass(workInactiveClass);
    $('#work-section-jobs').removeClass(workDefaultClass);
    $('#work-section-projects').removeClass(workDefaultClass);
    $('#work-section-jobs').scrollTop(0);
    $('#work-section-projects').scrollTop(0);
  }

  $.fn.selectSection = function() {
    clearWorkSelection();

    this.addClass(workSelectedClass)

    if (this[0].id == 'work-section-jobs') {
      $('#work-section-projects').addClass(workInactiveClass); 
    } else {
      $('#work-section-jobs').addClass(workInactiveClass); 
    }
  }

  animateSkills();

  $("#work-section-jobs.work-inactive h1, #work-section-jobs.work-default h1").on("click", function (){
    console.log('click')
    $("#work-section-jobs").selectSection(); 
  });
  $("#work-section-projects.work-inactive h1, #work-section-projects.work-default h1").on("click", function (){
    $("#work-section-projects").selectSection(); 
  });
});
