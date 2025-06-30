import { useNavigate } from "react-router-dom";
import SignupCard from "../components/features/signup/SignupCard";
import { useAuth } from "../context/AuthContext";
import { useEffect } from "react";

const Signup = () => {
  const navigate = useNavigate();
  const { isAuthenticated } = useAuth();

  useEffect(() => {
    if (isAuthenticated) {
      navigate('/dashboard');
    }
  }, [isAuthenticated, navigate]);

  return (
    <SignupCard />
  );
}

export default Signup;
