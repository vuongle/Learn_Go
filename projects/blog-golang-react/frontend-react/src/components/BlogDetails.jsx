import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axios from "axios";
import { REACT_APP_BACKEND_URL } from "../constants/constants.js";

const BlogDetail = () => {
  const navigate = useNavigate();
  const [singlePost, setSinglePost] = useState();
  const { id } = useParams();
  useEffect(() => {
    const User = localStorage.getItem("user");

    if (!User) {
      navigate("/login");
    }
  }, []);

  const singleBlog = () => {
    axios
      .get(
        `${REACT_APP_BACKEND_URL}/apis/posts/${id}`,
        {
          withCredentials: true,
        }
      )
      .then(function (response) {
        setSinglePost(response?.data?.data?.blog);
        console.log(response?.data?.data?.blog);
      })
      .catch(function (error) {
        console.log(error);
      })
      .then(function () {
      });
  };
  useEffect(() => {
    singleBlog();
  }, []);

  return (
    <div className="relative">
      <div className="max-w-3xl mb-10 rounded overflow-hidden flex flex-col mx-auto text-center">
        <div className="max-w-3xl mx-auto text-xl sm:text-4xl font-semibold hover:text-indigo-600 transition duration-500 ease-in-out inline-block mb-2">
          The Best Activewear from the Nordstrom Anniversary Sale
        </div>

        <img className="w-full h-96 my-4" src={singlePost?.image} alt="" />
        <p className="text-gray-700 text-base leading-8 max-w-2xl mx-auto">
          Author: {singlePost?.user?.first_name} {singlePost?.user?.last_name}
        </p>

        <hr />
      </div>

      <div className="max-w-3xl mx-auto">
        <div className="mt-3 bg-white rounded-b lg:rounded-b-none lg:rounded-r flex flex-col justify-between leading-normal">
          <div className="">
            <p className="text-base leading-8 my-5">{singlePost?.desc}</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default BlogDetail;
