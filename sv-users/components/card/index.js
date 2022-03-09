import * as React from "react";
import Card from "@mui/material/Card";

import CardContent from "@mui/material/CardContent";

export default function CardLayout({ children }) {
  return <Card sx={{ maxWidth: "100%" }}>{children}</Card>;
}
