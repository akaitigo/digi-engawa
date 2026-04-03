"use client";

interface FuriganaTextProps {
	text: string;
	furigana: string;
}

/**
 * ふりがなテキストを表示するコンポーネント。
 * furiganaが空の場合はtextをそのまま表示。
 * furiganaは「漢字(かんじ)」形式を <ruby> タグに変換する。
 */
export function FuriganaText({ text, furigana }: FuriganaTextProps) {
	if (!furigana) {
		return <span style={{ fontSize: "1.25rem", lineHeight: "2" }}>{text}</span>;
	}

	const parts = parseRuby(furigana);

	return (
		<span style={{ fontSize: "1.25rem", lineHeight: "2" }}>
			{parts.map((part, i) => {
				const key = `${part.text}-${String(i)}`;
				if (part.ruby) {
					return (
						<ruby key={key}>
							{part.text}
							<rp>(</rp>
							<rt style={{ fontSize: "0.75rem" }}>{part.ruby}</rt>
							<rp>)</rp>
						</ruby>
					);
				}
				return <span key={key}>{part.text}</span>;
			})}
		</span>
	);
}

interface RubyPart {
	text: string;
	ruby?: string;
}

function parseRuby(input: string): RubyPart[] {
	const parts: RubyPart[] = [];
	const regex = /([^\s(（]+)[（(]([^)）]+)[)）]/g;
	let lastIndex = 0;
	let match: RegExpExecArray | null = null;

	match = regex.exec(input);
	while (match !== null) {
		if (match.index > lastIndex) {
			parts.push({ text: input.slice(lastIndex, match.index) });
		}
		parts.push({ text: match[1] ?? "", ruby: match[2] });
		lastIndex = regex.lastIndex;
		match = regex.exec(input);
	}

	if (lastIndex < input.length) {
		parts.push({ text: input.slice(lastIndex) });
	}

	return parts;
}
