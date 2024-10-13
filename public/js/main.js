function showMenu() {
  let dropdown = document.getElementById("dropdown");
  dropdown.classList.toggle("hidden");
  dropdown.classList.toggle("flex");

  if (dropdown.classList.contains("flex")) {
    dropdown.classList.add("justify-center", "items-center");
    setTimeout(() => {
      document.addEventListener("click", closeMenu);
    }, 0);
  }

  function closeMenu(event) {
    if (!dropdown.contains(event.target)) {
      dropdown.classList.add("hidden");
      dropdown.classList.remove("flex", "justify-center", "items-center");
      document.removeEventListener("click", closeMenu);
    }
  }
}
