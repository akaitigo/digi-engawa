import type { Metadata, Viewport } from "next";
import type { ReactNode } from "react";

export const metadata: Metadata = {
	title: "デジタル縁側",
	description: "公民館・地域サロンで開催するデジタル教室の運営を支援するプラットフォーム",
	manifest: "/manifest.json",
};

export const viewport: Viewport = {
	width: "device-width",
	initialScale: 1,
	themeColor: "#4A7C59",
};

export default function RootLayout({ children }: { children: ReactNode }) {
	return (
		<html lang="ja">
			<body style={{ margin: 0, fontFamily: "sans-serif" }}>{children}</body>
		</html>
	);
}
