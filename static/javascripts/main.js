$(document).ready(function () {

  function animateSkills() {
    setTimeout(function () {
      if ($('#about-wrapper').length) {
        $('#about-wrapper').addClass('show-section');    
      }
    }, 500)
  }

  animateSkills();
});
