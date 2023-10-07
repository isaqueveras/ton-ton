import './globals.css'
import type { Metadata } from 'next'
import { IBM_Plex_Sans_Arabic } from 'next/font/google'

const font = IBM_Plex_Sans_Arabic({ 
  subsets: ['latin'],
  weight: ["100","200","400", "500", "600", "700"],
  fallback: ['arial']
})

export const metadata: Metadata = {
  title: 'Tón-Tón - Workspace',
  description: '',
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return <html lang="pt-BR"><body className={font.className}>{children}</body></html>
}
