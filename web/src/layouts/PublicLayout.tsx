import { Link, Outlet, useLocation } from 'react-router-dom';

const PublicLayout = () => {
  const location = useLocation();

  return (
    <div>
      <nav style={{ padding: '1rem' }}>
        <ul style={{ display: 'flex', gap: '1rem', listStyle: 'none', padding: 0 }}>
          <li>
            <Link to="/" style={{ fontWeight: location.pathname === '/' ? 'bold' : 'normal' }}>
              Home
            </Link>
          </li>
          <li>
            <Link to="/login" style={{ fontWeight: location.pathname === '/login' ? 'bold' : 'normal' }}>
              Login
            </Link>
          </li>
          <li>
            <Link to="/register" style={{ fontWeight: location.pathname === '/register' ? 'bold' : 'normal' }}>
              Register
            </Link>
          </li>
        </ul>
      </nav>

      <main>
        <Outlet />
      </main>
    </div>
  );
}

export default PublicLayout;
