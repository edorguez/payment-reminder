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
  isAuthenticated: boolean;
  signUp: (email: string, password: string, name: string) => Promise<void>;
  login: (email: string, password: string) => Promise<void>;
  loginGoogle: () => Promise<void>;
  logout: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | null>(null);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  useEffect(() => {
    const unsub = onAuthStateChanged(auth, (user: User | null) => {
      setIsAuthenticated(!!user);
    });
    return unsub;
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
    console.log('existing user', res);
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
      value={{ isAuthenticated, signUp, login, loginGoogle, logout }}
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
