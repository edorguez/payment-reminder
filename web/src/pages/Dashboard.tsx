import DashboardDetailCard from "../components/features/dashboard/DashboardDetailCard";
import DashboardTable from "../components/features/dashboard/DashboardTable";
import Container from "../components/ui/Container";

const Dashboard = () => {
  return (
    <Container>
      <div className="py-4 grid sm:grid-cols-3 gap-3">
        <DashboardDetailCard iconPath="/images/dashboard/bell.png" title="Upcoming Alerts" total={15} />
        <DashboardDetailCard iconPath="/images/dashboard/bell-sleep.png" title="Paused Alerts" total={5} />
        <DashboardDetailCard iconPath="/images/dashboard/list.png" title="Total Alerts" total={100} />
      </div>
      <DashboardTable />
    </Container>
  )
}

export default Dashboard;
