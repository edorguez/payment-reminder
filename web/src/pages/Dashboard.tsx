import DashboardDetailCard from "../components/features/dashboard/DashboardDetailCard";
import DashboardTable from "../components/features/dashboard/DashboardTable";
import Container from "../components/ui/Container";
import { HiBellAlert } from "react-icons/hi2";
import { FaClipboardList } from "react-icons/fa";
import { MdNotificationsPaused } from "react-icons/md";

const Dashboard = () => {
  return (
    <Container>
      <div className="py-4 grid sm:grid-cols-3 gap-3">
        <DashboardDetailCard icon={HiBellAlert} title="Upcoming Alerts" total={15} />
        <DashboardDetailCard icon={MdNotificationsPaused} title="Paused Alerts" total={5} />
        <DashboardDetailCard icon={FaClipboardList} title="Total Alerts" total={100} />
      </div>
      <DashboardTable />
    </Container>
  )
}

export default Dashboard;
