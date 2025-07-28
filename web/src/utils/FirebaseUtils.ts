import type { FirebaseError } from 'firebase/app';

export const GetFirebaseErrorMessage = (
    error: FirebaseError,
    customMsg: string,
): string => {
    switch (error.code) {
        case 'auth/email-already-in-use':
            return 'Email already in use';
        case 'auth/invalid-credential':
            return 'Email or password is invalid';
        case 'auth/popup-closed-by-user':
            return 'Please complete the login process';
        default:
            console.log(error.code);
            return customMsg || 'An error occurred';
    }
};
