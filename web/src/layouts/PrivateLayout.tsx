import { Link, Outlet, useLocation, useNavigate } from 'react-router-dom';

const PrivateLayout = () => {
  const location = useLocation();
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem('authToken');
    navigate('/login');
  };

  return (
    <div>
      <nav style={{ background: '#e0e0ff', padding: '1rem' }}>
        <ul style={{ display: 'flex', gap: '1rem', listStyle: 'none', padding: 0 }}>
          <li>
            <Link to="/dashboard" style={{ fontWeight: location.pathname === '/dashboard' ? 'bold' : 'normal' }}>
              Dashboard
            </Link>
          </li>
          <li>
            <Link to="/profile" style={{ fontWeight: location.pathname === '/profile' ? 'bold' : 'normal' }}>
              Profile
            </Link>
          </li>
          <li>
            <Link to="/settings" style={{ fontWeight: location.pathname === '/settings' ? 'bold' : 'normal' }}>
              Settings
            </Link>
          </li>
          <li>
            <button onClick={handleLogout}>Logout</button>
          </li>
        </ul>
      </nav>

      <main>
        <Outlet />
      </main>
    </div>
  );
}

export default PrivateLayout
