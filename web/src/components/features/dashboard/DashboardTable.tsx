import { FaPauseCircle } from "react-icons/fa";
import { MdEditNotifications, MdDelete } from "react-icons/md";

const DashboardTable = () => {
  return (
    <div className="overflow-x-auto bg-white rounded-md shadow-sm">
      <table className="table">
        <thead>
          <tr>
            <th></th>
            <th>Name</th>
            <th>Job</th>
            <th>Favorite Color</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr className="bg-base-200">
            <th>1</th>
            <td>Cy Ganderton</td>
            <td>Quality Control Specialist</td>
            <td>Blue</td>
            <td className="flex items-center justify-center gap-2">
              <button className="btn btn-sm btn-primary text-white">
                <FaPauseCircle />    
              </button>
              <button className="btn btn-sm btn-info text-white">
                <MdEditNotifications />
              </button>
              <button className="btn btn-sm btn-secondary text-white">
                <MdDelete />
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  )
};

export default DashboardTable;
