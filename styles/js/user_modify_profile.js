
// Colonne de gauche mdp 
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
// Colonne de droite changement mdp
const eye_change_confirmation = document.querySelector('#eye_change_password');
const eyeoff_change_confirmation = document.querySelector('#eye-off_change_password');
const passwordChangeFiel = document.querySelector('#password_change');
eye_change_confirmation.addEventListener('click', () => {
    eye_change_confirmation.style.display = "none";
    eyeoff_change_confirmation.style.display = "block";
    passwordChangeFiel.type = "text";
});
eyeoff_change_confirmation.addEventListener('click', () => {
    eyeoff_change_confirmation.style.display = "none";
    eye_change_confirmation.style.display = "block";
    passwordChangeFiel.type = "password";
});
// Colonne de droite confirmation changement mdp
const eye_confiramation_change_password = document.querySelector('#eye_confirmation_change_password');
const eyeoff_confirmation_change_password = document.querySelector('#eye-off_confirmation_change_password');
const passwordConfirmationChangeField = document.querySelector('#password_change_confirmation');
eye_confiramation_change_password.addEventListener('click', () => {
    eye_confiramation_change_password.style.display = "none";
    eyeoff_confirmation_change_password.style.display = "block";
    passwordConfirmationChangeField.type = "text";
});
eyeoff_confirmation_change_password.addEventListener('click', () => {
    eyeoff_confirmation_change_password.style.display = "none";
    eye_confiramation_change_password.style.display = "block";
    passwordConfirmationChangeField.type = "password";
});