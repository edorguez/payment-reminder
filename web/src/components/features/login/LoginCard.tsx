import { useNavigate } from "react-router-dom";
import { useAuth } from "../../../context/AuthContext";

const LoginCard = () => {
  const navigate = useNavigate();
  const { login } = useAuth();

  const handleLogin = () => {
    login('my-long-token');
    navigate('/dashboard');
  }

  return (
    <>
      <div className="flex justify-center mt-20">
        <fieldset className="fieldset bg-base-200 border-base-300 rounded-box w-xs border p-4">
          <label className="label">Email</label>
          <input type="email" className="input" placeholder="Email" />

          <label className="label">Password</label>
          <input type="password" className="input" placeholder="Password" />

          <button className="btn btn-neutral mt-4" onClick={handleLogin}>Login</button>
        </fieldset>
      </div>
    </>
  );
}

export default LoginCard;
