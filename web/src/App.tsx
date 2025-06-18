import { Link, useLocation, Outlet } from 'react-router-dom';

export default function App() {
  const location = useLocation();

  return (
    <div>
      <nav>
        <ul style={{ display: 'flex', gap: '1rem', listStyle: 'none', padding: 0 }}>
          <li>
            <Link 
              to="/" 
              style={{ fontWeight: location.pathname === '/' ? 'bold' : 'normal' }}
            >
              Home
            </Link>
          </li>
        </ul>
      </nav>

      <Outlet />
    </div>
  );
}
