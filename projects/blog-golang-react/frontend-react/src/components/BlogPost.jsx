import React, { useState, useEffect } from "react";
import axios from "axios";
import { Link } from "react-router-dom";
import { REACT_APP_BACKEND_URL } from "../constants/constants.js";

const BlogPost = () => {
  const [blogData, setBlogData] = useState([]);
  const [loading, setLoading] = useState(false);
  const allBlog = () => {
    setLoading(true);
    axios
      .get(
        `${REACT_APP_BACKEND_URL}/apis/posts`,
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
        console.log(error);
      });
  };
  useEffect(() => {
    allBlog();
  }, []);
  return (
    <>
      {loading && (
        <div className="text-2xl font-bold text-center px-56 pt-24">
          <h1>LOADING.....</h1>
        </div>
      )}
      <div className="container my-12 mx-auto px-4 md:px-12">
        <div className="flex flex-wrap -mx-1 lg:-mx-4">
          {blogData?.map((blog) => (
            <div key={blog?.id} className="my-1 px-1 w-full md:w-1/2 lg:my-4 lg:px-4 lg:w-1/3">
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
                  <a
                    className="no-underline text-grey-darker hover:text-red-dark"
                    href="#"
                  >
                    <span className="hidden">Like</span>
                    <i className="fa fa-heart"></i>
                  </a>
                </footer>
              </article>
            </div>
          ))}
        </div>
      </div>
    </>
  );
};
export default BlogPost;
