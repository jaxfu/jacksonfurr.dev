const spelldle = document.getElementById("spelldle");
const chatter = document.getElementById("chatter");
const timestrainer = document.getElementById("timestrainer");

const projects = [timestrainer, spelldle, chatter];
let count = 0;

function change_project(movement) {
	count += movement;
	const curr = Math.abs(count % projects.length);
	projects.forEach((p, i) => {
		p.style.display = curr === i % projects.length ? "block" : "none";
	});
}

async function send_message(e) {
	e.preventDefault();
	e.stopPropagation();

	const form = e.target;
	const email = form.email.value;
	const msg = form.message.value;

	const data = {
		contact_info: email,
		message: msg,
	};

	await fetch("/api/message", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(data),
	});

	form.innerHTML = "Message sent, thank you for reaching out!";
}

