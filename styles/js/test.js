const eye = document.querySelector('.feather-eye');
const eyeoff = document.querySelector('.feather-eye-off');
const passwordField = document.querySelector('input[type=password]');
eye.addEventListener('click', () => {
    eye.style.display = "none";
    eyeoff.style.display = "block";
    passwordField.type = "text";
});
eyeoff.addEventListener('click', () => {
    eyeoff.style.display = "none";
    eye.style.display = "block";
    passwordField.type = "password";
});
// Repeat the same logic for the confirmation field
const eye_confirmation = document.querySelector('#eye_confirmation');
const eyeoff_confirmation = document.querySelector('#eye-off_confirmation');
const passwordConfirmationField = document.querySelector('#password_confirmation');
eye_confirmation.addEventListener('click', () => {
    eye_confirmation.style.display = "none";
    eyeoff_confirmation.style.display = "block";
    passwordConfirmationField.type = "text";
});
eyeoff_confirmation.addEventListener('click', () => {
    eyeoff_confirmation.style.display = "none";
    eye_confirmation.style.display = "block";
    passwordConfirmationField.type = "password";
});

