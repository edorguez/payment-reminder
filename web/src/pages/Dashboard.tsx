import DashboardDetailCard from "../components/features/dashboard/DashboardDetailCard";
import DashboardTable from "../components/features/dashboard/DashboardTable";
import Container from "../components/ui/Container";

const Dashboard = () => {
  return (
    <Container>
      <div className="py-4 grid sm:grid-cols-3 gap-3">
        <DashboardDetailCard />
        <DashboardDetailCard />
        <DashboardDetailCard />
      </div>
      <DashboardTable />
    </Container>
  )
}

export default Dashboard;
