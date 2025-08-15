import {
  createContext,
  useContext,
  useEffect,
  useState,
  type ReactNode,
} from 'react';
import {
  createUserWithEmailAndPassword,
  updateProfile,
  type User,
} from 'firebase/auth';
import {
  auth,
  signInWithEmailAndPassword,
  signOut,
  onAuthStateChanged,
  signInWithPopup,
  googleProvider,
} from '../firebase';
import { accountService } from '../services/account/account.service';

interface AuthContextType {
  currentUser: User | null;
  loading: boolean;
  signUp: (email: string, password: string, name: string) => Promise<void>;
  login: (email: string, password: string) => Promise<void>;
  loginGoogle: () => Promise<void>;
  logout: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | null>(null);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [currentUser, setCurrentUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const unsubscribe = onAuthStateChanged(auth, async (user) => {
      if (user) {
        try {
          const idToken = await user.getIdToken();

          const userStatus = await accountService.checkUserStatus(idToken);

          if (userStatus.isAllowed) {
            // User is valid, set the current user and proceed
            setCurrentUser(user);
          } else {
            // User is not allowed (expired, banned, etc.).
            // Log them out of Firebase to prevent access.
            await signOut(auth);
            setCurrentUser(null);
            // You might want to show a message to the user here.
            console.error("User account is not active or has expired.");
          }
        } catch (error) {
          console.error("Error checking user status:", error);
          await signOut(auth);
          setCurrentUser(null);
        }
      } else {
        // No user is signed in.
        setCurrentUser(null);
      }
      setLoading(false);
    });

    return unsubscribe;
  }, []);

  const signUp = async (email: string, password: string, name: string) => {
    const cred = await createUserWithEmailAndPassword(
      auth,
      email,
      password,
    );
    await updateProfile(cred.user, { displayName: name });
  };

  const login = async (email: string, password: string) => {
    await signInWithEmailAndPassword(auth, email, password);
  };

  const loginGoogle = async () => {
    const userCredential = await signInWithPopup(auth, googleProvider);
    const userEmail = userCredential.user.email;

    const res = await accountService.getUser({ email: userEmail ?? '' });
    if (!res.ok) {
      if (res.status === 404) {
        await accountService.createUser();
      } else {
        console.error(res.status, res.message);
      }
    }
  };

  const logout = async () => {
    await signOut(auth);
  };

  return (
    <AuthContext.Provider
      value={{ currentUser, loading, signUp, login, loginGoogle, logout }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};
