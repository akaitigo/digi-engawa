"use client";

import { useCallback, useState } from "react";

interface ClassroomFormProps {
	onSubmit: (data: ClassroomFormData) => void;
	submitting?: boolean;
}

export interface ClassroomFormData {
	title: string;
	description: string;
	location: string;
	capacity: number;
	scheduled_at: string;
}

export function ClassroomForm({ onSubmit, submitting = false }: ClassroomFormProps) {
	const [title, setTitle] = useState("");
	const [description, setDescription] = useState("");
	const [location, setLocation] = useState("");
	const [capacity, setCapacity] = useState(20);
	const [date, setDate] = useState("");
	const [time, setTime] = useState("10:00");

	const handleSubmit = useCallback(
		(e: React.FormEvent) => {
			e.preventDefault();
			if (!title || !date) return;

			const scheduledAt = new Date(`${date}T${time}:00`).toISOString();
			onSubmit({ title, description, location, capacity, scheduled_at: scheduledAt });
		},
		[title, description, location, capacity, date, time, onSubmit],
	);

	const inputStyle = {
		width: "100%",
		padding: "12px",
		fontSize: "1.125rem",
		borderRadius: "8px",
		border: "2px solid #ccc",
		marginBottom: "1rem",
		boxSizing: "border-box" as const,
	};

	const labelStyle = {
		display: "block",
		fontSize: "1.125rem",
		fontWeight: "bold" as const,
		marginBottom: "0.5rem",
		color: "#333",
	};

	return (
		<form onSubmit={handleSubmit} style={{ maxWidth: "500px", margin: "0 auto" }}>
			<div>
				<label htmlFor="title" style={labelStyle}>
					教室名 *
				</label>
				<input
					id="title"
					type="text"
					value={title}
					onChange={(e) => setTitle(e.target.value)}
					required
					style={inputStyle}
				/>
			</div>

			<div>
				<label htmlFor="description" style={labelStyle}>
					説明
				</label>
				<textarea
					id="description"
					value={description}
					onChange={(e) => setDescription(e.target.value)}
					rows={3}
					style={{ ...inputStyle, resize: "vertical" }}
				/>
			</div>

			<div>
				<label htmlFor="location" style={labelStyle}>
					場所
				</label>
				<input
					id="location"
					type="text"
					value={location}
					onChange={(e) => setLocation(e.target.value)}
					style={inputStyle}
				/>
			</div>

			<div style={{ display: "flex", gap: "1rem" }}>
				<div style={{ flex: 1 }}>
					<label htmlFor="date" style={labelStyle}>
						日付 *
					</label>
					<input
						id="date"
						type="date"
						value={date}
						onChange={(e) => setDate(e.target.value)}
						required
						style={inputStyle}
					/>
				</div>
				<div style={{ flex: 1 }}>
					<label htmlFor="time" style={labelStyle}>
						時間
					</label>
					<input id="time" type="time" value={time} onChange={(e) => setTime(e.target.value)} style={inputStyle} />
				</div>
			</div>

			<div>
				<label htmlFor="capacity" style={labelStyle}>
					定員
				</label>
				<input
					id="capacity"
					type="number"
					min={1}
					max={100}
					value={capacity}
					onChange={(e) => setCapacity(Number(e.target.value))}
					style={inputStyle}
				/>
			</div>

			<button
				type="submit"
				disabled={submitting || !title || !date}
				style={{
					width: "100%",
					padding: "16px",
					fontSize: "1.25rem",
					fontWeight: "bold",
					minHeight: "56px",
					borderRadius: "12px",
					border: "none",
					backgroundColor: submitting || !title || !date ? "#ccc" : "#4A7C59",
					color: "#ffffff",
					cursor: submitting ? "wait" : "pointer",
					marginTop: "0.5rem",
				}}
			>
				{submitting ? "作成中..." : "教室を作成する"}
			</button>
		</form>
	);
}
