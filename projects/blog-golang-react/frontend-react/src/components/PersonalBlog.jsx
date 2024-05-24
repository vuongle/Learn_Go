import React, { useState, useEffect } from "react";
import axios from "axios";
import { useSnackbar } from "react-simple-snackbar";
import { useNavigate } from "react-router-dom";
import { REACT_APP_BACKEND_URL } from "../constants/constants.js";

const PersonalBlog = () => {
  const [blogData, setBlogData] = useState();
  const [loading, setLoading] = useState(false);
  const [deleteLoading, setDeleteLoading] = useState(false);
  const navigate = useNavigate();

  const options = {
    position: "bottom-right",
    style: {
      backgroundColor: "gray",
      border: "2px solid lightgreen",
      color: "white",
      fontFamily: "Menlo, monospace",
      fontSize: "20px",
      textAlign: "center",
    },
    closeStyle: {
      color: "lightcoral",
      fontSize: "16px",
    },
  };
  const [openSnackbar] = useSnackbar(options);
  const uniqueBlog = () => {
    setLoading(true);
    const User = JSON.parse(localStorage.getItem("user"));
    axios
      .get(
        `${REACT_APP_BACKEND_URL}/apis/posts-by-user/${User.id}`,
        {
          withCredentials: true,
        }
      )
      .then(function (response) {
        setLoading(false);
        setBlogData(response?.data?.data?.blogs);
        console.log(response?.data?.data?.blogs);
      })
      .catch(function (error) {
        setLoading(false);
      })
      .then(function () {
        // always executed
      });
  };

  useEffect(() => {
    const User = localStorage.getItem("user");
    if (!User) {
      navigate("/login");
    }
    uniqueBlog();
  }, []);

  const deleteBtn = (blog) => {
    setDeleteLoading(true);
    axios
      .delete(
        `${REACT_APP_BACKEND_URL}/apis/posts/${blog.id}`,

        {
          withCredentials: true,
        }
      )
      .then(function (response) {
        setDeleteLoading(false);
        openSnackbar("Delete blog successfully");

        uniqueBlog();

        console.log(response?.data);
      })
      .catch(function (error) {
        // handle error
        setDeleteLoading(false);
        openSnackbar(error?.response?.data?.message);
      })
      .then(function () {
        // always executed
      });
  };

  return (
    <>
      {!loading && blogData?.length <= 0 && (
        <div className="text-2xl font-bold text-center flex justify-center items-center pl-16 pt-24">
          <h1>You don't have post yet. Kindly create a post </h1>
        </div>
      )}
      {loading && (
        <div className="text-2xl font-bold text-center px-56 pt-24">
          <h1>LOADING.....</h1>
        </div>
      )}
      <div className="container my-12 mx-auto px-4 md:px-12">
        <div className="flex flex-wrap -mx-1 lg:-mx-4">
          {blogData?.map((blog) => (
            <div key={blog.id} className="my-1 px-1 w-full md:w-1/2 lg:my-4 lg:px-4 lg:w-1/3">
              <article className="overflow-hidden rounded-lg shadow-lg">
                <a href={`/detail/${blog.id}`}>
                  <img
                    alt="Placeholder"
                    className="block h-72 w-full"
                    src={blog?.image}
                  />
                </a>

                <header className="flex items-center justify-between leading-tight p-2 md:p-4">
                  <h1 className="text-lg">
                    <a
                      className="no-underline hover:underline text-black"
                      href={`/detail/${blog.id}`}
                    >
                      {blog.title}
                    </a>
                  </h1>
                  <p className="text-grey-darker text-sm">
                    {new Date(blog?.created_at).toLocaleString()}
                  </p>
                </header>

                <footer className="flex items-center justify-between leading-none p-2 md:p-4">
                  <a
                    className="flex items-center no-underline hover:underline text-black"
                    href={`/detail/${blog.id}`}
                  >
                    <img
                      alt="Placeholder"
                      className="block rounded-full w-5 h-5"
                      src={blog?.image}
                    />
                    <p className="ml-2 text-sm">
                      {blog?.user?.first_name} {blog?.user?.last_name}
                    </p>
                  </a>
                  <div>
                    <button
                      onClick={() => deleteBtn(blog)}
                      disabled={loading ? true : false}
                      className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
                    >
                      {deleteLoading ? "Loading" : "Delete"}
                    </button>
                  </div>
                  <div className="">
                    <a href={`edit/${blog.id}`}>
                      <button className="bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-2 px-4 rounded">
                        Edit
                      </button>
                    </a>
                  </div>
                </footer>
              </article>
            </div>
          ))}
        </div>
      </div>
    </>
  );
};

export default PersonalBlog;
