import Link from "next/link";

export default function Navbar() {
  return (
    <nav className="bg-white border-gray-200 dark:bg-gray-950 border-b-2">
      <div className="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
        <div className="flex items-center justify-center">
          <span className="self-center text-2xl font-semibold whitespace-nowrap dark:text-white mr-5"><Link href="/workspace">workspace</Link></span>
          <div className="hidden w-full md:flex md:w-auto md:order-1">
            <ul className="flex flex-col font-medium md:flex-row md:space-x-2 dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
              <li><Link href="/workspace/articles" className="py-1 px-3 hover:bg-yellow-500 hover:rounded hover:transition-colors text-gray-90 dark:text-white dark:hover:text-white">Articles</Link></li>
              <li><Link href="/workspace/new/article" className="py-1 px-3 hover:bg-yellow-500 hover:rounded hover:transition-colors text-gray-90 dark:text-white dark:hover:text-white">New Article</Link></li>
            </ul>
          </div>
        </div>

        <div className="flex items-center md:order-2">
          <button type="button" className="flex mr-3 text-sm bg-gray-800 rounded-full md:mr-0 focus:ring-4 focus:ring-gray-300 dark:focus:ring-gray-600" id="user-menu-button" aria-expanded="false" data-dropdown-toggle="user-dropdown" data-dropdown-placement="bottom">
            <span className="sr-only">Open user menu</span>
            <div className="relative inline-flex items-center justify-center w-8 h-8 overflow-hidden bg-gray-100 rounded-full dark:bg-gray-600">
              <span className="font-medium text-gray-600 dark:text-gray-300">IV</span>
            </div>
          </button>
          
          <div className="z-50 hidden my-4 text-base list-none bg-white divide-y divide-gray-100 rounded-lg shadow dark:bg-gray-700 dark:divide-gray-600" id="user-dropdown">
            <div className="px-4 py-3">
              <span className="block text-sm text-gray-900 dark:text-white">Isaque Veras</span>
              <span className="block text-sm  text-gray-500 truncate dark:text-gray-400">isaqueveras@tonton.io</span>
            </div>
            <ul className="py-2" aria-labelledby="user-menu-button">
              <li><a href="#" className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white">Workspace</a></li>
              <li><a href="#" className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white">Settings</a></li>
              <li><a href="#" className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white">Sign out</a></li>
            </ul>
          </div>
          
          <button data-collapse-toggle="navbar-user" type="button" className="inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600" aria-controls="navbar-user" aria-expanded="false">
            <span className="sr-only">Open main menu</span>
            <svg className="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 17 14">
              <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 1h15M1 7h15M1 13h15"/>
            </svg>
          </button>
        </div>
      </div>
    </nav>
  )
}