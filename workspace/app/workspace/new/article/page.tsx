'use client';

import { useState } from "react";
import Navbar from "@/components/navbar";
import Editor from "@/components/Editor";

export default function NewArticle() {
  const [data, setData] = useState<string>("");

  return (
    <>
      <Navbar/>
      <div className="max-w-screen-xl mx-auto p-4">
        <div className="grid grid-cols-12 gap-4">
          <div className="col-span-9 border-r-2 pr-3">
            <h1 className="text-xl font-medium">Title</h1>
            <input placeholder="Title" className="w-full border-2 p-3 text-gray-900 font-normal text-base" />
            <div>
              <h1 className="mt-4 text-xl font-medium">Content</h1>
              <Editor onChange={(data: string) => { setData(data) }} />
              <button type="submit" className="inline-flex px-5 py-2.5 text-base font-medium mt-4 text-center text-white bg-green-700 focus:ring-4 focus:ring-green-200 dark:focus:ring-green-900 hover:bg-green-900 hover:transition-all">Publish post</button>
            </div>
          </div>
          <div className="col-span-3">
            <div>
              <label id="listbox-label" className="block text-base font-medium leading-6 text-gray-900">Status</label>
            </div>
          </div>
        </div>
      </div>
    </>
  )
}