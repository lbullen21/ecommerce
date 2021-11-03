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
    const emailBlock = document.getElementById("email");
    const phoneBlock = document.getElementById("phone");

    const inputEmail = emailBlock.value;
    const inputPhone = phoneBlock.value;

    const email = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/;
    const phone = /^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/im

    const validEmail = inputEmail&&inputEmail.match(email)
    const validPhone = inputPhone&&inputPhone.match(phone)

    if(validEmail&&validPhone){
        alert("Thank you for your info!")
    } else if (!validEmail&&!validPhone){
        emailBlock.classList.add("not-valid")
        phoneBlock.classList.add("not-valid")
    } else if (!validEmail){
        emailBlock.classList.add("not-valid")
    } else if (!validPhone){
        phoneBlock.classList.add("not-valid")
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


//currency converter
function convertCurrency(){
    const apiKey = "793d7f349f7723249eab"
    const url = `https://free.currconv.com/api/v7/convert?q=USD_EUR&compact=ultra&apiKey=${apiKey}`

    $.get(url, function(data){
        const usdToEur = Object.values(data)[0]
        const dollarValue = document.getElementById("currency-converter").value
        const displayEuroValue = document.getElementById("currency-results")
        const euroValue = (dollarValue*usdToEur).toFixed(2)
        displayEuroValue.innerHTML = `Price in Euros: €${euroValue}`
        console.log(euroValue)
    })  

}

/*Display products on products page when tab clicked*/
function getProducts(){
   $.get("http://localhost:3000/products", function(data){
        console.log(data)
        displayProducts(JSON.parse(data));
        
    })


    console.log("working?")
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
                    <p class="price">$${product.price}  (${product.detail})</p> 
            </div>`
    
            $(".section__products").append(newHTML)

            
        })
    } 
     


/*Display search results*/
    function getSearch(){
        const searchTerm = document.getElementById("search-term").value
        const priceSort = document.getElementById("price-sort").value
        console.log(searchTerm)
        $.get("http://localhost:3000/search", { term: searchTerm, order: priceSort}, function(data){
            console.log(data)
            displaySearch(JSON.parse(data))
            
            // if (!searchTerm){
            //     alert("not valid")
            // }
        })
    }
        

    function displaySearch(productsJSON){
        $.each(productsJSON, function(i, product){
            $(".search_products").html("")
            newHTML = `<div class="product">
            <h3>${product.flavor}</h3>
                    <a href="./product_description.html"> <img 
                        class="product__img"
                        src=${product.photo}
                    /></a>
                    <p class="price">$${product.price}  (${product.detail})</p> 
            </div>`
    
            $(".search_products").append(newHTML)
            
        })
        
    }
    

