// assets/js/sidebar.js

function toggleSidebar() {
    const sidebar = document.getElementById("sidebar");
    const mainContent = document.getElementById("main-content");

    // Toggle the sidebar visibility
    sidebar.classList.toggle("-translate-x-full");

    // Move main content to the right when sidebar is visible
    mainContent.classList.toggle("ml-64");
}
