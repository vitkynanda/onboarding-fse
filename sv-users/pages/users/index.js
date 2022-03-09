import Layout from "../../components/layout";
import React from "react";
import UsersTable from "../../components/user-table";
import DialogLayout from "../../components/dialog";
import { TextField } from "@mui/material";
import AddIcon from "@mui/icons-material/Add";
import { useState } from "react";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { toast } from "react-toastify";
import { createNewUser, getListUsers } from "../../constants/service/api";

export default function Users() {
  const queryClient = useQueryClient();
  const [page, setPage] = useState(1);
  const { data, isLoading } = useQuery(["users", page], () =>
    getListUsers(page)
  );

  const [payload, setPayload] = React.useState({
    firstName: data?.data?.firstName,
    lastName: data?.data?.lastName,
    gender: "L",
    birthdate: "1995-12-20",
    active: true,
    contact: {
      email: "vitkynptr@gmail.com",
      phone: "0181123123",
    },
    hobbies: ["futsal", "mobile game", "programming"],
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setPayload({ ...payload, [name]: value });
  };

  const { mutate: addUser } = useMutation((data) => createNewUser(data), {
    onSuccess: (data) => {
      console.log(data);
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

  return (
    <Layout title="Users">
      <div className="flex items-center justify-between">
        <DialogLayout
          title="Add User"
          payload={payload}
          action={addUser}
          icons={<AddIcon />}
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
      </div>
      <UsersTable data={data} isLoading={isLoading} />
    </Layout>
  );
}

// export async function getServerSideProps(context) {
//   const users = await getListUsers();
//   console.log(users);
//   return {
//     props: {
//       data: users,
//     },
//   };
// }

// export async function getStaticProps(context) {
//   const users = await getListUsers();
//   console.log("exec");

//   return {
//     props: {
//       data: testFetch(),
//     },
//     revalidate: 1,
//   };
// }
