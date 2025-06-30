import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import LoginCard from "../components/features/login/LoginCard";
import { useAuth } from "../context/AuthContext";

const Login = () => {
  const navigate = useNavigate();
  const { isAuthenticated } = useAuth();

  useEffect(() => {
    if (isAuthenticated) {
      navigate('/dashboard');
    }
  }, [isAuthenticated, navigate]);

  return <LoginCard />;
};

export default Login;
