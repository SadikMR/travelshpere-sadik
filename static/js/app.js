/**
 * TravelSphere — app.js
 * Global UI behaviour: mobile navbar toggle.
 */
(function () {
  "use strict";

  var toggle = document.getElementById("nav-toggle");
  var links  = document.getElementById("nav-links");
  if (!toggle || !links) return;

  toggle.addEventListener("click", function () {
    var open = links.classList.toggle("open");
    toggle.setAttribute("aria-expanded", open ? "true" : "false");

    /* animate the three bars into an X */
    var bars = toggle.querySelectorAll("span");
    if (open) {
      bars[0].style.transform = "translateY(7px) rotate(45deg)";
      bars[1].style.opacity   = "0";
      bars[2].style.transform = "translateY(-7px) rotate(-45deg)";
    } else {
      bars[0].style.transform = "";
      bars[1].style.opacity   = "";
      bars[2].style.transform = "";
    }
  });

  /* close menu when a link is clicked (SPA-style nav feel) */
  links.querySelectorAll(".nav-link").forEach(function (a) {
    a.addEventListener("click", function () {
      links.classList.remove("open");
      toggle.setAttribute("aria-expanded", "false");
    });
  });

}());