import DashboardIcon from "@mui/icons-material/Dashboard";
import GroupIcon from "@mui/icons-material/Group";

const routes = [
  {
    name: "Dashboard",
    path: "/",
    isVisited: false,
    icon: <DashboardIcon />,
  },
  {
    name: "Users",
    path: "/users",
    isVisited: false,
    icon: <GroupIcon />,
  },
];

export default routes;
