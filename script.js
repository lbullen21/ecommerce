//hamburger menu
const hamburger = document.querySelector(".hamburger");
const navMenu = document.querySelector(".nav-menu");

hamburger.addEventListener("click", mobileMenu);

function mobileMenu() {
    hamburger.classList.toggle("active");
    navMenu.classList.toggle("active");
}

const navLink = document.querySelectorAll(".nav-link");

navLink.forEach(n => n.addEventListener("click", closeMenu));

function closeMenu() {
    hamburger.classList.remove("active");
    navMenu.classList.remove("active");
}

//form validation


//slideshow functionality
let slideIndex = 1
showSlides(slideIndex)

//shows previous/next slide
function plusSlides(n){
    showSlides(slideIndex += n)
}
//shows current slide
function currentSlide(n){
    showSlides(slideIndex = n)
}

function showSlides(n){
    let i;
    let slides = document.getElementsByClassName("mySlides")
    let dots = document.getElementsByClassName("dot")

    if(n>slideIndex){
        slideIndex = 1
    } else if(n < 1){
        slideIndex = slides.length
    }

    for(i = 0; i<slides.length; i++){
        slides[i].style.display = "none"
    }

    for(i = 0; i<dots.length ; i++){
        dots[i].className = dots[i].className.replace("active", " ")
    }
    
    slides[slideIndex - 1].style.display = "block";
    dots[slideIndex - 1].className += "  active"

}