"use client";

import { useCallback, useEffect, useRef, useState } from "react";
import { API_BASE, WS_BASE } from "@/lib/api";
import type { HelpRequest } from "@/types/help";
import { HelpRequestCard } from "./HelpRequestCard";

interface HelpBoardProps {
	classroomId: string;
}

const MAX_RECONNECT_DELAY = 30000;

export function HelpBoard({ classroomId }: HelpBoardProps) {
	const [requests, setRequests] = useState<HelpRequest[]>([]);
	const [connected, setConnected] = useState(false);
	const reconnectDelay = useRef(1000);

	useEffect(() => {
		const loadInitial = async () => {
			try {
				const res = await fetch(`${API_BASE}/api/classrooms/${classroomId}/help-requests`);
				if (res.ok) {
					const data = (await res.json()) as HelpRequest[];
					setRequests(data);
				}
			} catch {
				// Initial load failure is non-critical
			}
		};
		void loadInitial();
	}, [classroomId]);

	useEffect(() => {
		let socket: WebSocket | null = null;
		let reconnectTimer: ReturnType<typeof setTimeout> | null = null;
		let unmounted = false;

		const connect = () => {
			if (unmounted) return;

			try {
				socket = new WebSocket(`${WS_BASE}/classroom/${classroomId}`);

				socket.onopen = () => {
					if (!unmounted) {
						setConnected(true);
						reconnectDelay.current = 1000;
					}
				};

				socket.onclose = () => {
					if (!unmounted) {
						setConnected(false);
						reconnectTimer = setTimeout(() => {
							reconnectDelay.current = Math.min(reconnectDelay.current * 2, MAX_RECONNECT_DELAY);
							connect();
						}, reconnectDelay.current);
					}
				};

				socket.onerror = () => {
					socket?.close();
				};

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
		};

		connect();

		return () => {
			unmounted = true;
			if (reconnectTimer !== null) clearTimeout(reconnectTimer);
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
			// Network failure
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
					role="status"
					aria-label={connected ? "サーバーに接続中" : "サーバーに未接続"}
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
							onStatusChange={(reqId, status) => void handleStatusChange(reqId, status)}
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
							onStatusChange={(reqId, status) => void handleStatusChange(reqId, status)}
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
