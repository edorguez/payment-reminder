import {
  createContext,
  useContext,
  useEffect,
  useState,
  type ReactNode,
} from 'react';
import { useAuth } from './AuthContext';
import type { ProfileContextType } from '../types/profile';
import { accountService } from '../services/account/account.service';
import type { UserDto } from '../services/account/account.types';
import { notify } from '../lib/toast';

const ProfileContext = createContext<ProfileContextType | null>(null);

export const ProfileProvider = ({ children }: { children: ReactNode }) => {
  const { currentUser, loading: authLoading } = useAuth();
  const [user, setUser] = useState<UserDto | null>(null);
  const [loading, setLoading] = useState(true);

  const fetchProfile = async () => {
    if (!currentUser) {
      setUser(null);
      setLoading(false);
      return;
    }

    setLoading(true);
    try {
      const res = await accountService.getUser({ email: currentUser.email ?? '' });

      if (res.ok) {
        setUser(res.data);
      } else {
        notify.error(res.message);
        setUser(null);
      }
    } catch (error) {
      console.error("Failed to fetch user profile:", error);
      notify.error("Failed to fetch user profile");
      setUser(null);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (!authLoading) {
      fetchProfile();
    }
  }, [currentUser, authLoading]);

  const value = {
    user,
    loading,
    refetchProfile: fetchProfile,
  };

  return (
    <ProfileContext.Provider value={value}>
      {children}
    </ProfileContext.Provider>
  );
};

export const useProfile = () => {
  const context = useContext(ProfileContext);
  if (!context) {
    throw new Error('useProfile must be used within a ProfileProvider');
  }
  return context;
};
