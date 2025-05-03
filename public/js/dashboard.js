document.addEventListener('DOMContentLoaded', function() {
    // Dark mode functionality
    initDarkMode();
    
    // Navigation functionality
    initNavigation();
});

/**
 * Initialize dark mode toggle functionality
 */
function initDarkMode() {
    const darkModeToggle = document.getElementById('darkModeToggle');
    const lightIcon = document.getElementById('lightIcon');
    const darkIcon = document.getElementById('darkIcon');
    
    // Check for saved user preference, if any
    const isDarkMode = localStorage.getItem('darkMode') === 'true';
    
    // Set initial dark mode state
    if (isDarkMode) {
        document.documentElement.classList.add('dark');
        lightIcon.classList.remove('hidden');
        darkIcon.classList.add('hidden');
    }
    
    // Toggle dark mode on button click
    darkModeToggle.addEventListener('click', function() {
        const isDarkModeActive = document.documentElement.classList.contains('dark');
        
        if (isDarkModeActive) {
            // Switch to light mode
            document.documentElement.classList.remove('dark');
            darkIcon.classList.remove('hidden');
            lightIcon.classList.add('hidden');
            localStorage.setItem('darkMode', 'false');
        } else {
            // Switch to dark mode
            document.documentElement.classList.add('dark');
            lightIcon.classList.remove('hidden');
            darkIcon.classList.add('hidden');
            localStorage.setItem('darkMode', 'true');
        }
    });
}

/**
 * Initialize navigation functionality
 */
function initNavigation() {
    // Get all navigation links
    const navLinks = document.querySelectorAll('.nav-link');
    
    // Get all content sections
    const contentSections = document.querySelectorAll('.content-section');
    
    // Add click event listener to each nav link
    navLinks.forEach(link => {
        link.addEventListener('click', function(e) {
            e.preventDefault();
            
            // Get the target section ID from the href attribute
            const targetId = this.getAttribute('href').substring(1);
            const targetContentId = `${targetId}-content`;
            
            // Remove active class from all links
            navLinks.forEach(navLink => {
                navLink.classList.remove('active', 'bg-gray-100', 'dark:bg-gray-700');
                navLink.classList.add('hover:bg-gray-100', 'dark:hover:bg-gray-700');
            });
            
            // Add active class to clicked link
            this.classList.add('active', 'bg-gray-100', 'dark:bg-gray-700');
            this.classList.remove('hover:bg-gray-100', 'dark:hover:bg-gray-700');
            
            // Hide all content sections
            contentSections.forEach(section => {
                section.classList.add('hidden');
            });
            
            // Show the target content section
            const targetContent = document.getElementById(targetContentId);
            if (targetContent) {
                targetContent.classList.remove('hidden');
            }
            
            // Special case for logout
            if (targetId === 'logout') {
                handleLogout();
            }
        });
    });
}

/**
 * Handle logout functionality
 */
function handleLogout() {
    // Clear any stored authentication tokens
    localStorage.removeItem('authToken');
    localStorage.removeItem('userRole');
    
    // Redirect to login page
    // In a real application, you'd also want to call a logout API endpoint
    alert('You have been logged out. Redirecting to login page...');
    // window.location.href = '/login.html';
}

/**
 * Fetch data from API with authentication token
 * @param {string} endpoint - API endpoint to fetch data from
 * @returns {Promise} - Promise resolving to JSON response
 */
async function fetchApi(endpoint) {
    const token = localStorage.getItem('authToken');
    
    try {
        const response = await fetch(endpoint, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        
        if (!response.ok) {
            throw new Error(`API request failed: ${response.status}`);
        }
        
        return await response.json();
    } catch (error) {
        console.error('API fetch error:', error);
        throw error;
    }
}