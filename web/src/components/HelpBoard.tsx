"use client";

import { useCallback, useEffect, useState } from "react";
import type { HelpRequest } from "@/types/help";
import { HelpRequestCard } from "./HelpRequestCard";

interface HelpBoardProps {
	classroomId: string;
}

const API_BASE = process.env["NEXT_PUBLIC_API_URL"] ?? "http://localhost:8080";
const WS_BASE = process.env["NEXT_PUBLIC_WS_URL"] ?? "ws://localhost:8080/ws";

export function HelpBoard({ classroomId }: HelpBoardProps) {
	const [requests, setRequests] = useState<HelpRequest[]>([]);
	const [connected, setConnected] = useState(false);

	useEffect(() => {
		const loadInitial = async () => {
			try {
				const res = await fetch(`${API_BASE}/api/classrooms/${classroomId}/help-requests`);
				if (res.ok) {
					const data = (await res.json()) as HelpRequest[];
					setRequests(data);
				}
			} catch {
				// Silently fail for initial load
			}
		};
		void loadInitial();
	}, [classroomId]);

	useEffect(() => {
		const wsUrl = `${WS_BASE}/classroom/${classroomId}`;
		let socket: WebSocket | null = null;

		try {
			socket = new WebSocket(wsUrl);

			socket.onopen = () => setConnected(true);
			socket.onclose = () => setConnected(false);
			socket.onerror = () => setConnected(false);

			socket.onmessage = (event) => {
				try {
					const msg = JSON.parse(String(event.data)) as { type: string; data: HelpRequest };
					if (msg.type === "help_request_created") {
						setRequests((prev) => [...prev, msg.data]);
					} else if (msg.type === "help_request_updated") {
						setRequests((prev) => prev.map((r) => (r.id === msg.data.id ? msg.data : r)));
					}
				} catch {
					// Ignore malformed messages
				}
			};
		} catch {
			// WebSocket not available
		}

		return () => {
			socket?.close();
		};
	}, [classroomId]);

	const handleStatusChange = useCallback(async (id: string, status: string) => {
		try {
			const res = await fetch(`${API_BASE}/api/help-requests/${id}`, {
				method: "PATCH",
				headers: { "Content-Type": "application/json" },
				body: JSON.stringify({ status }),
			});

			if (res.ok) {
				const updated = (await res.json()) as HelpRequest;
				setRequests((prev) => prev.map((r) => (r.id === updated.id ? updated : r)));
			}
		} catch {
			// Silently fail
		}
	}, []);

	const pendingRequests = requests.filter((r) => r.status === "pending");
	const activeRequests = requests.filter((r) => r.status === "in_progress");
	const resolvedRequests = requests.filter((r) => r.status === "resolved");

	return (
		<div style={{ maxWidth: "600px", margin: "0 auto", padding: "1rem" }}>
			<div
				style={{
					display: "flex",
					justifyContent: "space-between",
					alignItems: "center",
					marginBottom: "1rem",
				}}
			>
				<h2 style={{ fontSize: "1.5rem", fontWeight: "bold" }}>ヘルプボード</h2>
				<span
					style={{
						fontSize: "0.875rem",
						padding: "4px 8px",
						borderRadius: "12px",
						backgroundColor: connected ? "#e8f5e9" : "#ffebee",
						color: connected ? "#2e7d32" : "#c62828",
					}}
				>
					{connected ? "● 接続中" : "○ 未接続"}
				</span>
			</div>

			{pendingRequests.length > 0 && (
				<section style={{ marginBottom: "1.5rem" }}>
					<h3 style={{ fontSize: "1.25rem", marginBottom: "0.5rem" }}>
						🔴 たすけて ({String(pendingRequests.length)})
					</h3>
					{pendingRequests.map((r) => (
						<HelpRequestCard
							key={r.id}
							request={r}
							onStatusChange={(id, status) => void handleStatusChange(id, status)}
						/>
					))}
				</section>
			)}

			{activeRequests.length > 0 && (
				<section style={{ marginBottom: "1.5rem" }}>
					<h3 style={{ fontSize: "1.25rem", marginBottom: "0.5rem" }}>
						🟡 たいおうちゅう ({String(activeRequests.length)})
					</h3>
					{activeRequests.map((r) => (
						<HelpRequestCard
							key={r.id}
							request={r}
							onStatusChange={(id, status) => void handleStatusChange(id, status)}
						/>
					))}
				</section>
			)}

			{resolvedRequests.length > 0 && (
				<section>
					<h3 style={{ fontSize: "1.25rem", marginBottom: "0.5rem" }}>
						🟢 かいけつ ({String(resolvedRequests.length)})
					</h3>
					{resolvedRequests.map((r) => (
						<HelpRequestCard key={r.id} request={r} />
					))}
				</section>
			)}

			{requests.length === 0 && (
				<p style={{ textAlign: "center", fontSize: "1.125rem", color: "#999", padding: "2rem" }}>
					まだヘルプリクエストはありません
				</p>
			)}
		</div>
	);
}
