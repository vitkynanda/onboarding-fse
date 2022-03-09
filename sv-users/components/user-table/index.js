/* eslint-disable @next/next/no-img-element */
import * as React from "react";
import { styled } from "@mui/material/styles";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell, { tableCellClasses } from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import {
  getListUsers,
  editUserData,
  deleteUserData,
} from "../../constants/service/api";
import { useQuery, useMutation, useQueryClient } from "react-query";
import DialogLayout from "../dialog";
import EditIcon from "@mui/icons-material/Edit";
import DeleteIcon from "@mui/icons-material/Delete";
import { toast } from "react-toastify";
import TextField from "@mui/material/TextField";
import { CircularProgress } from "@mui/material";

const StyledTableCell = styled(TableCell)(({ theme }) => ({
  [`&.${tableCellClasses.head}`]: {
    backgroundColor: "rgb(59 130 246 / 0.5)",
    color: theme.palette.common.white,
  },
  [`&.${tableCellClasses.body}`]: {
    fontSize: 14,
  },
}));

const StyledTableRow = styled(TableRow)(({ theme }) => ({
  "&:nth-of-type(odd)": {
    backgroundColor: theme.palette.action.hover,
  },
  // hide last border
  "&:last-child td, &:last-child th": {
    border: 0,
  },
}));

const Row = ({ row }) => {
  const queryClient = useQueryClient();
  const [payload, setPayload] = React.useState({
    firstName: row.firstName,
    lastName: row.lastName,
    ...row,
  });

  const { mutate: editUser } = useMutation((data) => editUserData(data), {
    onSuccess: (data) => {
      if (data.status_code === 200) {
        queryClient.invalidateQueries("users");
        toast.success(data.message);
      } else {
        toast.error(data.message);
      }
    },
    onError: (err) => {
      toast.error(err.message);
    },
  });

  const { mutate: deleteUser } = useMutation((data) => deleteUserData(data), {
    onSuccess: (data) => {
      if (data.status_code === 200) {
        queryClient.invalidateQueries("users");
        toast.success(data.message);
      } else {
        toast.error(data.message);
      }
    },
    onError: (err) => {
      toast.error(err.message);
    },
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setPayload({ ...payload, [name]: value });
  };

  return (
    <StyledTableRow>
      <StyledTableCell component="th" scope="row">
        <p className="capitalize">{row.firstName}</p>
      </StyledTableCell>
      <StyledTableCell>
        <p className="capitalize">{row.lastName}</p>
      </StyledTableCell>
      <StyledTableCell>
        <p className="capitalize">{row.birthdate}</p>
      </StyledTableCell>
      <StyledTableCell>
        <p className="capitalize">{row.gender}</p>
      </StyledTableCell>
      <StyledTableCell>
        <p className="capitalize">{row.active ? "active" : "inactive"}</p>
      </StyledTableCell>
      <StyledTableCell>
        <div className="flex space-x-2">
          <DialogLayout
            title="Edit User"
            data={row}
            payload={payload}
            action={editUser}
            icons={<EditIcon />}
          >
            <div className="space-y-3 w-72 py-2">
              <TextField
                id="outlined-basic"
                label="First Name"
                name="firstName"
                value={payload.firstName}
                onChange={handleChange}
                variant="outlined"
                className="w-full"
              />
              <TextField
                id="outlined-basic"
                label="Last Name"
                name="lastName"
                value={payload.lastName}
                onChange={handleChange}
                variant="outlined"
                className="w-full"
              />
            </div>
          </DialogLayout>
          <DialogLayout
            title="Delete User"
            action={deleteUser}
            data={row}
            icons={<DeleteIcon />}
          >
            <div>Are you sure want to delete this user ?</div>
          </DialogLayout>
        </div>
      </StyledTableCell>
    </StyledTableRow>
  );
};

const UsersTable = ({ data, isLoading }) => {
  React.useEffect(() => {}, []);

  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 700 }} aria-label="customized table">
        <TableHead>
          <TableRow>
            <StyledTableCell>First Name</StyledTableCell>
            <StyledTableCell>Last Name</StyledTableCell>
            <StyledTableCell>Birth Date</StyledTableCell>
            <StyledTableCell>Gender</StyledTableCell>
            <StyledTableCell>Status</StyledTableCell>
            <StyledTableCell>Action</StyledTableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {isLoading ? (
            <StyledTableRow>
              <StyledTableCell colSpan={6}>
                <div className="w-full flex items-center justify-center h-72">
                  <CircularProgress />
                </div>
              </StyledTableCell>
            </StyledTableRow>
          ) : (
            data?.data?.map((row, id) => <Row key={id} row={row} />)
          )}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default UsersTable;
