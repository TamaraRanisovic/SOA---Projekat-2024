document.getElementById("loginForm").addEventListener("submit", function(event) {
    event.preventDefault(); // Prevent form submission

    var formData = new FormData(this); // Get form data
    var username = formData.get("username");
    var password = formData.get("password");

    // Send login credentials to Go backend
    fetch("/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ username: username, password: password })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error("Failed to login");
        }
        return response.json();
    })
    .then(data => {
        // Store JWT token in localStorage
        localStorage.setItem("jwtToken", data.token);
        // Redirect to dashboard or another page
        window.location.href = "/login.html"; // Change this to the desired URL
    })
    .catch(error => {
        document.getElementById("message").textContent = "Failed to login";
        console.error("Error:", error);
    });
});
