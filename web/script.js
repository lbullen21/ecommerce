/* Hamburger Menu */
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

/* Form Validation */

function checkValid() {

    const inputEmail = document.getElementById("email").value;
    const inputPhone = document.getElementById("phone").value;

    const email = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/;
    const phone = /^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/im

    const validEmail = inputEmail&&inputEmail.match(email)
    const validPhone = inputPhone&&inputPhone.match(phone)

    if(validEmail&&validPhone){
        alert("Thank you for your info!")
    } else if (!validEmail){
        alert("Please add a valid email address.")
    } else if (!validPhone){
        alert("Please add a valid phone number.")
    }
}

$(".form-submit").click(function(){
    checkValid();
    event.preventDefault();
})




/* Slideshow functionality on homepage */
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

    for (i = 0; i < slides.length; i++) {
        slides[i].style.display = "none";  
      }
      slideIndex++;
      if (slideIndex > slides.length) {slideIndex = 1}    
      for (i = 0; i < dots.length; i++) {
        dots[i].className = dots[i].className.replace(" active", "");
      }
      slides[slideIndex-1].style.display = "block";  
      dots[slideIndex-1].className += " active";
      setTimeout(showSlides, 4000); // Change image every 2 seconds

}

/*Display products on products page when tab clicked*/
$(".product-nav").click(function(){
    function getProducts(){
        $.get("http://localhost:3000/prducts", function(data){
            displayProducts(JSON.parse(data))
        })
    }
    
    //prepares product html
    function displayProducts(productsJSON){
        $.each(productsJSON, function(i, product){
            newHTML = `<div class="product">
            <h3>${product.flavor}</h3>
                    <a href="./product_description.html"> <img 
                        class="product__img"
                        src=${product.photo}
                    /></a>
                    <p class="price">${product.price} ${product.detail}</p> 
            </div>`
    
            $(".section_products").append(newHTML)
        })
    }
})    
